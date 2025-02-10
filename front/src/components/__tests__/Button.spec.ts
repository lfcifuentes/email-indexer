import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Button from '../Button.vue'

describe('Button.vue', () => {
  it('renders default slot content if no label prop is provided', () => {
    const wrapper = mount(Button, {
      slots: { default: 'Search' }
    });
    expect(wrapper.text()).toContain('Search');
  });

  it('renders label prop if provided and no slot content is given', () => {
    const wrapper = mount(Button, {
      props: { label: 'Click Me' }
    });
    // Se renderiza el fallback del slot, que usará label en ausencia de contenido en el slot.
    expect(wrapper.text()).toContain('Click Me');
  });

  it('renders slot content overriding the label prop', () => {
    const wrapper = mount(Button, {
      props: { label: 'Click Me' },
      slots: { default: 'Override' }
    });
    expect(wrapper.text()).toContain('Override');
    expect(wrapper.text()).not.toContain('Click Me');
  });

  it('renders with default type "submit" when not provided', () => {
    const wrapper = mount(Button, {
      slots: { default: 'Click' }
    });
    const button = wrapper.find('button');
    expect(button.attributes('type')).toBe('submit');
  });

  it('renders with provided type attribute', () => {
    const wrapper = mount(Button, {
      props: { type: 'button' },
      slots: { default: 'Click' }
    });
    const button = wrapper.find('button');
    expect(button.attributes('type')).toBe('button');
  });
});
