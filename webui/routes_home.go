package webui

import (
	"context"
	"net/http"

	"github.com/delaneyj/gomponents-iconify/iconify/vscode_icons"
	. "github.com/delaneyj/toolbelt/gomps"
	"github.com/go-chi/chi/v5"
)

func setupHome(ctx context.Context, router *chi.Mux) error {

	type Feature struct {
		Description string
		Icon        NODE
		Details     NODE
	}

	// features := []Feature{
	// 	{
	// 		Description: "Fine Grained Reactivity via Signals",
	// 		Icon:        ph.GitDiff(),
	// 		Details:     DIV(TXT("No Virtual DOM. proxy wrappers, or re-rendering the entire page on every change.  Take the best available options and use hassle free.")),
	// 	},
	// 	{
	// 		Description: "Declarative Batteries Included (but optional)",
	// 		Icon:        game_icons.Batteries(),
	// 		Details: DIV(
	// 			CLS("breadcrumbs"),
	// 			UL(
	// 				CLS(
	// 					"flex flex-wrap gap-2 justify-center items-center",
	// 				),
	// 				LI(TXT("Custom Actions")),
	// 				LI(TXT("Attribute Binding")),
	// 				LI(TXT("Focus")),
	// 				LI(TXT("Signals")),
	// 				LI(TXT("DOM Events")),
	// 				LI(TXT("Refs")),
	// 				LI(TXT("Intersects")),
	// 				LI(TXT("Two-Way Binding")),
	// 				LI(TXT("Visibility")),
	// 				LI(TXT("Teleporting")),
	// 				LI(TXT("Text Replacement")),
	// 				LI(TXT("HTMX like features")),
	// 				LI(TXT("Server Sent Events")),
	// 				LI(TXT("Redirects")),
	// 				LI(TXT("View Transition API")),
	// 				LI(TXT("BigInt Support")),
	// 			),
	// 		),
	// 	},
	// }

	languages := []NodeFunc{
		vscode_icons.FileTypeAssembly,
		vscode_icons.FileTypeApl,
		vscode_icons.FileTypeC,
		vscode_icons.FileTypeCpp,
		vscode_icons.FileTypeClojure,
		vscode_icons.FileTypeCsharp,
		vscode_icons.FileTypeGoGopher,
		vscode_icons.FileTypeJava,
		vscode_icons.FileTypeJs,
		vscode_icons.FileTypeKotlin,
		vscode_icons.FileTypeLua,
		vscode_icons.FileTypePython,
		vscode_icons.FileTypeRust,
		vscode_icons.FileTypeTypescript,
		vscode_icons.FileTypeZig,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, Page(
			DIV(
				CLS("container my-5 mx-auto"),
				DIV(
					CLS("grid grid-cols-1 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-3 gap-16"),
					RANGE(languages, func(fn NodeFunc) NODE {
						return DIV(
							CLS("card h-full w-full shadow-2xl ring-4 bg-base-300 ring-secondary text-secondary-content"),
							DIV(
								CLS("card-body flex flex-col justify-center items-center"),
								DIV(
									CLS("avatar avatar-xl"),
									fn(
										CLS("w-16 h-12 md:w-24 md:h-24 mask bg-gradient-to-t from-base-200 to-base-300 p-2 md:p-4 mask-hexagon"),
									),
								),
							),
						)
					}),
				),
			),
		))
	})

	return nil
}

/**

DIV(
							CLS("card w-full shadow-2xl ring-4 bg-base-300 ring-secondary text-secondary-content"),
							DIV(
								CLS("card-body flex flex-col justify-center items-center"),
								UL(
									CLS("flex flex-col gap-6 justify-center items-center text-2xl gap-4  max-w-xl"),
									RANGE(features, func(f Feature) NODE {
										return LI(
											DIV(
												CLS("flex flex-col gap-1 justify-center items-center"),
												DIV(
													CLS("flex gap-2 items-center"),
													f.Icon,
													TXT(f.Description),
												),
												DIV(
													CLS("text-lg opacity-50 p-2 rounded"),

													f.Details,
												),
											),
										)
									}),
								),
							),
						),**/

/**DIV(
	CLS("flex flex-col gap-4 items-center"),
	DIV(
		CLS("flex flex-wrap gap-1 md:gap-2 justify-center items-center text-6xl"),
		RANGE(languages, func(fn NodeFunc) NODE {
			return DIV(
				CLS("avatar avatar-xl"),
				fn(
					CLS("w-16 h-12 md:w-24 md:h-24 mask bg-gradient-to-t from-base-200 to-base-300 p-2 md:p-4 mask-hexagon"),
				),
			)
		}),
	),
),*/
