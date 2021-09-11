import { shallowMount } from "@vue/test-utils";
import App from "@/App.vue";

describe("App.vue", () => {
  it("should title is not null", () => {
    const msg = "ToDo List";
    const wrapper = shallowMount(App);
    expect(wrapper.text()).toMatch(msg);
  });

  it("should list is valid", () => {
    const wrapper = shallowMount(App);
    wrapper.setData({ tasks: [{ id: 1, task: "task1" }] });
    expect(wrapper.classes("li"));
  });

  it("should error is valid", () => {
    const wrapper = shallowMount(App);
    wrapper.setData({ error: true });
    expect(wrapper.classes("error"));
  });

  it("should input is valid", () => {
    const wrapper = shallowMount(App);
    expect(wrapper.classes("input"));
  });

  it("should button is valid", () => {
    const wrapper = shallowMount(App);
    expect(wrapper.classes("button"));
  });
});
