const utils = require('./utils');

test('float_random_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('float_random_default = 0');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('float_random_default = 7');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
