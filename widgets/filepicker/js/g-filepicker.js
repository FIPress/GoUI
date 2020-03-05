class GoUIFilePicker extends HTMLElement {
    static get observedAttributes() {
        return ['accept','multiple', 'editable','style','class'];
    }

    constructor() {
        super();

        this.settings = {
            //isSave: false,
            //message: "",
            //fileTypes: "",
            //startLocation:"",
            //suggestedFilename:"",
            //multiple:false,
            //fileOnly:false,
            //dirOnly:false,
            //allowsOtherFileTypes:false,
            //canCreateDir:false,
            //showsHiddenFiles:false,
        };
    }

    get editable() {
        return !this.editor.getAttribute("readonly");
    }

    set editable(v) {
        const isEditable = Boolean(v);
        this.editor.setAttribute("readonly",!isEditable);
    }

    set accept(v) {
        //this.fileBtn.setAttribute("accept",v);
        this.settings.accept = v;
    }

    set multiple(v) {
        const isMulti = Boolean(v);
        //this.fileBtn.setAttribute("multiple",v);
        this.settings.multiple = isMulti;
    }

    set action(v) {
        this.settings.isSave = v == "save";
    }

    get value() {
        //return this.fileBtn.value;
        return this.editor.value;
    }

    set value(v) {
        this.editor.text = v;
    }

    get files() {
        //split
        return this.editor.value;
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

        let _this=this;

        this.btn.onclick = function() {
            //console.log("request goui service");
            goui.request({url:"filepicker",
                data:_this.settings,
                success:function(path) {
                    _this.editor.value = path;
                }});
        };


        if(this.hasAttribute("style")) {
            el.setAttribute("style",this.getAttribute("style"));
        }

        if(this.hasAttribute("class")) {
            //const s = getComputedStyle(this);
            el.setAttribute("class",this.getAttribute("class"));
        }

        if(this.hasAttribute("multiple")) {
            //_this.fileBtn.setAttribute("multiple",this.getAttribute("multiple"));
            const isMulti = Boolean(this.getAttribute("multiple"));
            //this.fileBtn.setAttribute("multiple",v);
            _this.settings.multiple = isMulti;
        }

        if(this.hasAttribute("accept")) {
            //_this.fileBtn.setAttribute("accept",this.getAttribute("accept"));
            _this.settings.accept = this.getAttribute("accept");
        }

        if(this.hasAttribute("action")) {
            _this.settings.isSave = this.getAttribute("action") == "save";
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
customElements.define('g-filepicker', GoUIFilePicker);