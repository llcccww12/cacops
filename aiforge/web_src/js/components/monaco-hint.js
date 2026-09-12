
function getSuggestions(monaco , model, position, keywords, snippts) {
  let word = model.getWordUntilPosition(position)
  let range = {
    startLineNumber: position.lineNumber,
    endLineNumber: position.lineNumber,
    startColumn: word.startColumn,
    endColumn: word.endColumn
  }
  let rs = keywords.map(item => {
    return {
      label: item,
      kind: monaco.languages.CompletionItemKind.Keyword,
      insertText: item,
      insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
      range
    }
  })

  snippts.map(item => {
    rs.push({
      ...item,
      range
    })
  })

  return rs
}



export default (monaco) => {
  const ifelse = {
    label: 'ifelse',
    kind: monaco.languages.CompletionItemKind.Method,
    insertText: [
      'if (${1:condition}) {',
      '\t$0',
      '} else {',
      '\t',
      '}'
    ].join('\n'),
    insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet
  }

  const cKeywords = ['auto', 'break', 'case', 'char', 'const', 'continue', 'default', 'do', 'double', 'else', 'enum', 'extern',
    'float', 'for', 'goto', 'if', 'int', 'long', 'register', 'return', 'short', 'signed', 'sizeof', 'static', 'struct',
    'switch', 'typedef', 'union', 'unsigned', 'void', 'volatile', 'while', 'inline', 'restrict', '_Bool', '_Complex',
    '_Imaginary', '_Alignas', '_Alignof', '_Atomic', '_Static_assert', '_Noreturn', '_Thread_local', '_Generic']

  monaco.languages.registerCompletionItemProvider('cpp', {
    provideCompletionItems: (model, position) => {
      let suggestions = getSuggestions(monaco, model, position, cKeywords, [ifelse])
      return { suggestions }
    }
  })

  const pythonkeywords = ['False', 'None', 'True', 'and', 'as', 'assert', 'break', 'class', 'continue', 'def', 'del', 'elif',
    'else', 'except', 'finally', 'for', 'from', 'global', 'if', 'import', 'in', 'is', 'lambda', 'nonlocal', 'not', 'or',
    'pass', 'raise', 'return', 'try', 'while', 'with', 'yield']

  monaco.languages.registerCompletionItemProvider('python', {
    provideCompletionItems: (model, position) => {
      let snippets = [{
        label: 'print',
        kind: monaco.languages.CompletionItemKind.Snippet,
        insertText: [
          'print($0)',
        ].join('\n'),
        insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
      }]
      let suggestions = getSuggestions(monaco, model, position, pythonkeywords, snippets)
      return { suggestions }
    }
  })

  const javaKeywords = ['abstract', 'assert', 'boolean', 'break', 'byte', 'case', 'catch', 'char', 'class', 'const',
    'continue', 'default', 'do', 'double', 'else', 'enum', 'extends', 'final', 'finally', 'float', 'for', 'goto', 'if',
    'implements', 'import', 'instance of', 'int', 'interface', 'long', 'native',
    'new', 'package', 'private', 'protected', 'public', 'return', 'strictfp', 'short', 'static', 'super', 'switch',
    'synchronized', 'this', 'throw', 'throws', 'transient', 'try', 'void', 'volatile', 'while']

  monaco.languages.registerCompletionItemProvider('java', {
    provideCompletionItems: (model, position) => {
      let snippets = [ifelse,
        {
          label: 'main',
          kind: monaco.languages.CompletionItemKind.Snippet,
          insertText: [
            'public static void main(String[] args) {',
            '\t$0',
            '}',
          ].join('\n'),
          insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
        },
        {
          label: 'System.out.print',
          kind: monaco.languages.CompletionItemKind.Snippet,
          insertText: [
            'System.out.print($0)',
          ].join('\n'),
          insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
        }
      ]
      let suggestions = getSuggestions(monaco, model, position, javaKeywords, snippets)
      return { suggestions }
    }
  })

}

export const tipTxt = '该任务关卡设置了禁止复制粘贴，请手动输入代码。'