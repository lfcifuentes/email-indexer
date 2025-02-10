import { mount } from "@vue/test-utils"
import { describe, expect, it } from "vitest"
import SearchMessage from "../SearchMessage.vue"

describe('SearchMessage.vue', () => {
  it('renders correctly with default menu items', () => {
    const wrapper = mount(SearchMessage)
    // Verifica que se renderice el título y algunos elementos del menú.
    expect(wrapper.text()).toContain('Search Email')
    expect(wrapper.text()).toContain('Please')
    expect(wrapper.text()).toContain('email')
  })
})
