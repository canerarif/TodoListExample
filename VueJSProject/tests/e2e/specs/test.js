// For authoring Nightwatch tests, see
// https://nightwatchjs.org/guide

module.exports = {
  "default e2e tests": (browser) => {
    browser
      .init()
      .waitForElementVisible("#app")
      .assert.containsText("h1", "ToDo List")
      .assert.elementCount("img", 0)
      .waitForElementVisible("input")
      .waitForElementVisible("button")
      .click('button')
      .waitForElementVisible(".error")
      .end();
  },
};
