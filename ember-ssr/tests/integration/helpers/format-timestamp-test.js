import { module, test } from 'qunit';
import { setupRenderingTest } from 'ember-qunit';
import { render } from '@ember/test-helpers';
import hbs from 'htmlbars-inline-precompile';

module('Integration | Helper | format-timestamp', function (hooks) {
  setupRenderingTest(hooks);

  // Replace this with your real tests.
  test('it renders', async function (assert) {
    this.set('d', '07-29-2025');

    await render(hbs`{{format-timestamp this.d}}`);

    assert.equal(this.element.textContent.trim(), 'Jul 29, 2025 00:00.00 AM');
  });
});
