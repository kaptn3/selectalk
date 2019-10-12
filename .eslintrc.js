module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    '@nuxtjs',
    'plugin:nuxt/recommended'
  ],
  // add your custom rules here
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'indent': ['error', 2],
    'vue/script-indent': ['error', 2, { 'baseIndent': 1 }],
    'quote-props': 0,
    'comma-dangle': 0,
    'import/no-unresolved': 0,
    'vue/html-closing-bracket-spacing': 0,
    'no-plusplus': ['error', { 'allowForLoopAfterthoughts': true }],
    'no-unused-expressions': ['error', { 'allowTernary': true }],
    'prefer-destructuring': ['error', { 'array': false, 'object': true }],
    'max-len': ['error', {
      'code': 100,
      'ignoreTemplateLiterals': true,
      'ignoreStrings': true,
      'ignoreRegExpLiterals': true,
      'ignoreUrls': true,
      'ignorePattern': '^<path[\\s\\S]*\\/>$'
    }],
    'vue-a11y/click-events-have-key-events': 0,
    'space-before-function-paren': 0,
    'semi': 0
  },
  overrides: [
    {
      files: ['*.vue'],
      rules: {
        'indent': 'off'
      }
    }
  ]
}
