const utils = require('./utils');

test('bool_known_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('bool_known_default = true');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('bool_known_default = false');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
