const utils = require('./utils');

test('bool_empty_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('bool_empty_default = false');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('bool_empty_default = true');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
