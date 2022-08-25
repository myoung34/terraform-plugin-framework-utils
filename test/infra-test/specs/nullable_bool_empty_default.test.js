const utils = require('./utils');

test('nullable_bool_empty_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('nullable_bool_empty_default = false');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('nullable_bool_empty_default = true');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
