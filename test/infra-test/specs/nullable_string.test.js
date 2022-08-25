const utils = require('./utils');

test('nullable_string', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('nullable_string = ""');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('nullable_string = "One"');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
