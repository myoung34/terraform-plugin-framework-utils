const utils = require('./utils');

test('string_random_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('string_random_default = ""');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('string_random_default = "Three"');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
