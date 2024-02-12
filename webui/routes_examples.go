package webui

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/delaneyj/datastar"
	"github.com/delaneyj/gomponents-iconify/iconify/material_symbols"
	"github.com/delaneyj/toolbelt"
	. "github.com/delaneyj/toolbelt/gomps"
	"github.com/go-chi/chi/v5"
	"github.com/go-sanitize/sanitize"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/anchor"
	"mvdan.cc/xurls/v2"
)

var sanitizer *sanitize.Sanitizer

var markdownConverter = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.NewLinkify(
			extension.WithLinkifyAllowedProtocols([][]byte{
				[]byte("http:"),
				[]byte("https:"),
			}),
			extension.WithLinkifyURLRegexp(
				xurls.Strict(),
			),
		),
		emoji.Emoji,
		&anchor.Extender{},
		highlighting.NewHighlighting(
			highlighting.WithStyle("gruvbox"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
			),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
		parser.WithAttribute(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithXHTML(),
		html.WithUnsafe(),
	),
)

func examplePage(w http.ResponseWriter, r *http.Request, children ...NODE) error {
	nameParts := strings.Split(r.URL.Path, "/")
	name := nameParts[len(nameParts)-1]
	markdownPath := fmt.Sprintf("static/examples/%s.md", name)

	mdBytes, err := staticFS.ReadFile(markdownPath)
	if err != nil {
		return fmt.Errorf("error reading examples dir: %w", err)
	}

	mdBuf := bytes.NewBuffer(nil)
	if err := markdownConverter.Convert(mdBytes, mdBuf); err != nil {
		return fmt.Errorf("error converting markdown: %w", err)
	}

	back := A(
		HREF("/examples"),
		CLS("btn btn-primary"),
		material_symbols.ArrowBack(),
		TXT("Back to Examples"),
	)

	Render(w, Page(
		DIV(
			CLS("flex flex-col items-center p-8 gap-8"),
			back,
			DIV(
				CLS("flex flex-col max-w-5xl w-full prose"),
				RAW(mdBuf.String()),
				GRP(children...),
			),
			back,
		),
	))

	return nil
}

func setupExamples(ctx context.Context, router *chi.Mux) (err error) {
	sanitizer, err = sanitize.New()
	if err != nil {
		return fmt.Errorf("error creating sanitizer: %w", err)
	}
	return Route(ctx, router, "/examples", func(ctx context.Context, examplesRouter chi.Router) error {
		type Example struct {
			Pattern     string
			Description string
		}
		type ExampleGroup struct {
			Label    string
			Examples []Example
		}
		examples := []ExampleGroup{
			{
				Label: "Ported HTMX Examples*",
				Examples: []Example{
					{Pattern: "Click to Edit", Description: "Demonstrates inline editing of a data object"},
					{Pattern: "Bulk Update", Description: "Demonstrates bulk updating of multiple rows of data"},
					{Pattern: "Animations", Description: "Demonstrates Animations"},
				},
			},
		}

		examplesRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			Render(w, Page(
				DIV(
					CLS("flex flex-col items-center p-16"),
					DIV(
						CLS("flex flex-col gap-8 max-w-5xl"),
						RANGE(examples, func(g ExampleGroup) NODE {
							return DIV(
								DIV(
									DIV(
										CLS("text-4xl font-bold text-primary"),
										TXT(g.Label+"*"),
									),
									HR(
										CLS("divider border-primary"),
									),
								),
								TABLE(
									CLS("table w-full"),
									THEAD(TR(
										TH(TXT("Pattern")),
										TH(TXT("Description")),
									)),
									TBODY(
										RANGE(g.Examples, func(e Example) NODE {
											return TR(
												CLS("hover"),
												TD(A(
													CLS("link-secondary disable"),
													HREF("/examples/"+toolbelt.Cased(e.Pattern, toolbelt.Snake, toolbelt.Lower)),
													TXT(e.Pattern),
												)),
												TD(
													CLS("text-sm"),
													TXT(e.Description),
												),
											)
										}),
									),
								),
							)
						}),
						DIV(
							CLS("text-accent font-bold italic"),
							TXT("All examples use server-side logic in Go but you can use any language you like."),
						),
					),
				),
			))
		})

		if err := errors.Join(
			setupExamplesClickToEdit(ctx, examplesRouter),
			setupExamplesBulkUpdate(ctx, examplesRouter),
			setupExamplesAnimations(ctx, examplesRouter),
		); err != nil {
			return fmt.Errorf("error setting up examples routes: %w", err)
		}

		return nil
	})
}

var SignalStore = GRP(
	H4(TXT("Signal Store")),
	PRE(datastar.Text("ctx.JSONStringify(ctx.store())")),
)
