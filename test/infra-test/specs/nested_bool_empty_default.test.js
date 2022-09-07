const utils = require('./utils');

test('nested_bool_empty_default', () => {
  utils.reset();

  utils.writeConfig('');
  utils.testApply();

  utils.testApply();

  utils.writeConfig('  nested = {\n  }');
  utils.testApply();

  utils.writeConfig('  nested = {\n    bool_empty_default = false\n  }');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.writeConfig('  nested = {\n    bool_empty_default = true\n  }');
  utils.testApply();

  utils.writeConfig('');
  utils.testApply();

  utils.testDestroy();
});
