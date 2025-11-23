/* Quantum Element! Element is defined in code (html) and immediately collapses into standard html when code is executed */

"use strict";

class BUI_NAVBAR extends HTMLElement {
	static observedAttributes = [];

	constructor() {
		super();
		const shadow = this.attachShadow({ mode: "open" });

		const style = document.createElement("style");
		style.textContent = css;
		shadow.appendChild(style);

        console.log("section rendered:", this.innerHTML);
		const slot = document.createElement("slot");
        shadow.appendChild(slot)
	}

	connectedCallback() {
		this._internals = this.attachInternals();
	}
}

class BUI_NAVBAR_SECTION extends HTMLElement {
	static observedAttributes = [];

	constructor() {
		super();
	}

	connectedCallback() {
		this._internals = this.attachInternals();
	}

	render() {
		return this.innerHTML;
	}
}

customElements.define("bui-navbar", BUI_NAVBAR);
customElements.define("bui-navbar-section", BUI_NAVBAR_SECTION);
