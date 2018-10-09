export default class Upload {
    constructor () {
        this.element = document.querySelector('.js-upload');
        this.event = new Event('reload');

        this._dragenter = this._dragenter.bind(this);
        this._dragleave = this._dragleave.bind(this);
        this._dragover = this._dragover.bind(this);
        this._drop = this._drop.bind(this);

        this.element.addEventListener('dragenter', this._dragenter);
        this.element.addEventListener('dragover', this._dragover);
        this.element.addEventListener('dragleave', this._dragleave);
        this.element.addEventListener('drop', this._drop);
    }

    _dragenter () {
        this.element.classList.add('hover');
    }

    _dragover (event) {
        this.element.classList.add('hover');
        event.preventDefault();
    }

    _dragleave () {
        this.element.classList.remove('hover');
    }

    _drop (event) {
        event.preventDefault();

        let files = event.dataTransfer.files;

        let promises = [];

        ([...files]).forEach(file => {
            promises.push(this._upload(file));
        })

        Promise.all(promises).then(() => this.element.classList.remove('hover'))
            .then(() => this.element.dispatchEvent(this.event));

    }

    _upload (file) {
        let form = new FormData;
        form.append('file', file);

        let opts = {
            method: 'POST',
            body: form
        }

        return fetch('/upload', opts)
    }
}