module.exports = {
  root: true,
  env: {
    node: true
  },
  'extends': [
    'plugin:vue/essential',
    '@vue/standard'
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    "quotes": ["error", "single"],
    "semi": ["error", "always"],
    "no-empty": 2,
    "no-eq-null": 2,
    "no-empty-pattern":0,
    "no-new": 0,
    "no-fallthrough": 0,
    "no-unreachable": 0,
    "comma-dangle": 0,
    "vue/no-parsing-error": [0, { "x-invalid-end-tag": false }]
  },
  parserOptions: {
    parser: 'babel-eslint'
  }

}