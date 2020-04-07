class GoUIFilePicker extends HTMLElement {
    static get observedAttributes() {
        return ['accept','multiple', 'editable','style','class'];
    }

    bindAttributes() {
        this.attributesMap = new Map();
        this.attributesMap.set('action',{dataType:"string"});
        this.attributesMap.set('message',{dataType:"string"});
        this.attributesMap.set('accept',{dataType:"string"});
        this.attributesMap.set('multiple',{dataType: "bool"});
        this.attributesMap.set('dirOnly',{dataType: "bool"});
        this.attributesMap.set('fileOnly',{dataType: "bool"});
        this.attributesMap.set('editable',{el:this.$editor, set:function (v) {
                const readonly = !v;
                this.editor.setAttribute("readonly",readonly.toString());
            }});
        this.attributesMap.set('class',{el:this.$wrapper,set:function(v) {
                this.$wrapper.addClass(v);
            } });


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

        const shadow = this.attachShadow({mode: 'open'});
        this.$wrapper = $('<span>');

        this.$editor = $('<input type="text">');
        this.$wrapper.append(this.$editor);

        this.$btn = $('<button>Browse...</button>');
        this.$wrapper.append(this.$btn);
        shadow.appendChild(this.$wrapper[0]);

        this.bindAttributes();

        let _this=this;

        this.$btn.click(function() {
            goui.request({url:"filepicker",
                data:_this.data,
                success:function(path) {
                    _this.$editor.val(path);
                }});
        });
    }

    get value() {
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
        if(this.attributesMap) {
            let _this = this;
            this.attributesMap.forEach(function (v,k) {
                if(_this.hasAttribute(k)) {
                    if(v.set) {
                        v.set(_this.getAttribute(k));
                    }
                }
            });
        }
    }

    disconnectedCallback() {
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if(this.attributesMap) {
            let v = this.attributesMap.get(name);
            if(v && v.set) {
                var obj = this;
                if(v.el) {
                    obj = v.el;
                }
                v.set.call(obj,newVal);
            }
        }
    }

}

// Define the new element
customElements.define('g-filepicker', GoUIFilePicker);