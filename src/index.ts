import { LitElement, html } from "lit";
import { customElement, property, } from "lit/decorators.js";

@customElement("custom-hello")
export class Hello extends LitElement {
  @property()
  greetings: string = "World";

  protected render(): unknown {
    return html`
        <link rel="stylesheet" href="/styles/style.css">
        <h1>Hello ${this.greetings}!</h1>
      `
  }
}
