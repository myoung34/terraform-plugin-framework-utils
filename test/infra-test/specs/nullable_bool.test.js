const utils = require('./utils');

test('nullable_bool', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('nullable_bool = false');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('nullable_bool = true');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
