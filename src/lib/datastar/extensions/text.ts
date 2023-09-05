import { NamespacedReactiveRecords, functionGenerator } from '..'
import { addDataExtension } from '../core'

export const TEXT = Symbol('text')
export function addTextDataExtension() {
  addDataExtension(TEXT, {
    withExpression: ({ name, el, expression, dataStack, actions, reactivity: { effect } }) => {
      const signalFn = functionGenerator(expression)

      const elementData: NamespacedReactiveRecords = {
        text: {
          [name]: effect(() => {
            if (!dataStack?.signals) return
            const res = signalFn(el, dataStack, actions)
            el.textContent = `${res}`
          }),
        },
      }

      return elementData
    },
  })
}