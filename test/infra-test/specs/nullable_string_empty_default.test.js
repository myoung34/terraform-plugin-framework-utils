const utils = require('./utils');

test('nullable_string_empty_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('nullable_string_empty_default = ""');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('nullable_string_empty_default = "Three"');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
