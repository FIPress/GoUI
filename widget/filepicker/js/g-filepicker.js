class GoUIFilePicker extends HTMLElement {
    static get observedAttributes() {
        return ['accept','multiple', 'editable','canCreateDir','style','class'];
    }

    bindAttributes() {
        this.attributesMap = new Map();
        this.attributesMap.set('action',{dataType:"string"});
        this.attributesMap.set('message',{dataType:"string"});
        this.attributesMap.set('accept',{dataType:"string"});
        this.attributesMap.set('multiple',{dataType: "bool"});
        this.attributesMap.set('dirOnly',{dataType: "bool"});
        this.attributesMap.set('fileOnly',{dataType: "bool"});
        this.attributesMap.set('canCreateDir',{dataType: "bool"});
        this.attributesMap.set('editable',{set:function (v) {
                const readonly = !v;
                this.editor.setAttribute("readonly",readonly.toString());
            }});
        this.attributesMap.set('class',{set:function(v) {
                this.wrapper.classList.add(v);
            } });
        this.attributesMap.set('compact',{set:function(v) {
                if(!v) {
                    this.btn.innerText = "Browser...";
                } else {
                    this.btn.innerText = "...";
                }

            } });

        this.attributesMap.set('value',{set:function (v) {
                this.editor.value = v;
            },get:function (v) {
                return this.editor.value ;
            }});

        this.attributesMap.set('files',{set:function (v) {
                this.editor.value = v;
            },get:function (v) {
                return this.editor.value ;
            }});

        let _this = this;
        this.attributesMap.forEach(function(v,k) {
            Object.defineProperty(_this,k,v);
            if(v.dataType) {
                v.set = function (val) {
                    switch (v.dataType) {
                        case "bool":
                            _this.data[k] = !!val;
                            break;
                        default :
                            _this.data[k] = val;
                    }
                }
            }

        });
    }

    constructor() {
        super();

        this.data = {};
        this.inited = false;
    }

    connectedCallback() {
        if(!this.inited) {
            this.wrapper = document.createElement('span');
            this.wrapper.classList.add('g-filepicker');

            this.editor = document.createElement('input');
            this.editor.setAttribute('type', 'text');
            this.wrapper.appendChild(this.editor);

            this.btn = document.createElement('button');
            this.btn.innerText = 'Browse...';
            this.wrapper.appendChild(this.btn);

            const style = document.createElement('style');

            style.textContent = `
      .g-filepicker {
      width: 100%;
      display:grid;
      grid-template-columns: auto min-content;
      grid-gap: 0; 
    }
    `;

            // Attach the created elements to the shadow dom
            this.appendChild(style);
            //console.log(style.isConnected);
            this.appendChild(this.wrapper);

            this.bindAttributes();
            this.getAttributesFromDom();
            this.bindEvents();




            this.inited = true;
        }
    }

    disconnectedCallback() {
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if(this.attributesMap) {
            let v = this.attributesMap.get(name);
            this.invokeSet(v,newValue);
        }
    }

    getAttributesFromDom() {
        if(this.attributesMap) {
            let _this = this;
            this.attributesMap.forEach(function (v,k) {
                if(_this.hasAttribute(k)) {
                    _this.invokeSet(v,_this.getAttribute(k));
                }
            });
        }
    }

    bindEvents() {
        //elem.addEventListener('build', function (e) { /* ... */ }, false);

        let _this=this;

        this.btn.addEventListener('click',function() {
            goui.request({url:"filepicker",
                data:_this.data,
                success:function(path) {
                    _this.editor.value = path;
                    let event = new CustomEvent('pick',{"detail":path});
                    _this.dispatchEvent(event);
                }});
        });
    }

    invokeSet(attr,val) {
        if(!attr || !attr.set) {
            return;
        }

        attr.set.call(this,val);
    }

}

// Define the new element
customElements.define('g-filepicker', GoUIFilePicker);