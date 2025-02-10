import { describe, it, expect } from 'vitest'

import { mount } from '@vue/test-utils'
import Navbar from '../Navbar.vue'

describe('Navbar.vue', () => {
  it('renders correctly with default menu items', () => {
    const wrapper = mount(Navbar)
    // Verifica que se renderice el título y algunos elementos del menú.
    expect(wrapper.text()).toContain('Email Indexer')
    expect(wrapper.text()).toContain('Search')
    expect(wrapper.text()).toContain('Users')
  })

  it('toggles dark mode when clicking the theme toggle button', async () => {
    // Asegúrate de borrar la clase dark del html
    document.documentElement.classList.remove('dark')
    const wrapper = mount(Navbar)
    const themeButton = wrapper.find('#theme-toggle')
    // Simula el click para activar dark mode
    await themeButton.trigger('click')
    expect(document.documentElement.classList.contains('dark')).toBe(true)
    // Vuelve a hacer click para desactivar dark mode
    await themeButton.trigger('click')
    expect(document.documentElement.classList.contains('dark')).toBe(false)
  })
})
