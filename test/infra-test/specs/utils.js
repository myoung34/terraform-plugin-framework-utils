const { writeFileSync, readFileSync, unlinkSync, existsSync } = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const clean = (output) => {
  return output.replace(
    /test_team.default: (Creation|Modifications|Destruction) complete after \d+s/,
    '[CLEANED]',
  );
};

const statePath = path.resolve(__dirname, '..', 'terraform.tfstate');

module.exports = {
  reset() {
    if (existsSync(statePath)) {
      unlinkSync(statePath);
    }
  },
  writeConfig(config) {
    const configPath = path.resolve(__dirname, '..', 'main.tf');
    writeFileSync(
      configPath,
      `resource "test_team" "default" {
  id = "default"
  ${config}
}`,
    );
  },
  testApply() {
    const stdout = clean(
      execSync('terraform apply -auto-approve').toString('utf-8'),
    );
    expect(stdout).toMatchSnapshot();

    const state = JSON.parse(readFileSync(statePath, 'utf-8')).resources;
    expect(state).toMatchSnapshot();
  },
  testDestroy() {
    const stdout = clean(
      execSync('terraform destroy -auto-approve').toString('utf-8'),
    );
    expect(stdout).toMatchSnapshot();
  },
};
