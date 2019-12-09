class GoUIFilePicker extends HTMLElement {
    static get observedAttributes() {
        return ['accept','multiple', 'editable','style','class'];
    }

    constructor() {
        super();
    }

    get editable() {
        return !this.editor.getAttribute("readonly");
    }

    set editable(v) {
        const isEditable = Boolean(v);
        this.editor.setAttribute("readonly",!isEditable);
    }

    set accept(v) {
        this.fileBtn.setAttribute("accept",v);
    }

    set multiple(v) {
        const isMulti = Boolean(v);
        this.fileBtn.setAttribute("multiple",v);
    }

    get value() {
        return this.fileBtn.value;
    }

    set value(v) {
        this.editor.text = v;
    }

    get files() {
        return this.fileBtn.files;
    }

    connectedCallback() {
        const el = document.createElement('div');
        const shadow = el.attachShadow({mode: 'open'});
        const wrapper = document.createElement('span');

        this.editor = document.createElement('input');
        this.editor.setAttribute('type', 'text');
        wrapper.appendChild(this.editor);

        this.btn = document.createElement('button');
        this.btn.innerText = 'Browse...';
        wrapper.appendChild(this.btn);

        //this.fileBtn = document.createElement('input');
        //this.fileBtn.setAttribute('type', 'file');
        //this.fileBtn.setAttribute('style','opacity:0');
        //wrapper.appendChild(this.fileBtn);

        // Create some CSS to apply to the shadow dom
        const style = document.createElement('style');

        style.textContent = `
      span {
      	display:grid;
      	grid-template-columns: auto min-content;
      	grid-gap: 0; 
      }
    `;

        // Attach the created elements to the shadow dom
        shadow.appendChild(style);
        //console.log(style.isConnected);
        shadow.appendChild(wrapper);
        this.appendChild(el);

        var _this=this;

        this.btn.onclick = function() {
            console.log("request goui service");
        }


        if(this.hasAttribute("style")) {
            el.setAttribute("style",this.getAttribute("style"));
        }

        if(this.hasAttribute("class")) {
            //const s = getComputedStyle(this);
            el.setAttribute("class",this.getAttribute("class"));
        }

        if(this.hasAttribute("multiple")) {
            _this.fileBtn.setAttribute("multiple",this.getAttribute("multiple"));
        }

        if(this.hasAttribute("accept")) {
            _this.fileBtn.setAttribute("accept",this.getAttribute("accept"));
        }

        if(this.hasAttribute("capture")) {
            _this.fileBtn.setAttribute("capture",this.getAttribute("capture"));
        }

        const readonly = !this.getAttribute("editable");
        this.editor.setAttribute("readonly",readonly);

    }

    disconnectedCallback() {
    }

    attributeChangedCallback(name, oldValue, newValue) {
        switch(name) {

        }
    }


}

// Define the new element
customElements.define('g-filePicker', GoUIFilePicker);