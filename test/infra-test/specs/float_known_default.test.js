const utils = require('./utils');

test('float_known_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('float_known_default = 1');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('float_known_default = 3');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('float_known_default = 0');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
