const utils = require('./utils');

test('string_known_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('string_known_default = "One"');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('string_known_default = "Three"');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('string_known_default = ""');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
