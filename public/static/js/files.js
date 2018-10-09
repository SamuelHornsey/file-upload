export default class Files {
    constructor () {
        this.element = document.querySelector('.js-files');
        this.upload = document.querySelector('.js-upload');
        this.control = document.querySelector('.js-grid-control');

        this._get = this._get.bind(this);
        this._toggleGrid = this._toggleGrid.bind(this);

        this._get();
        this.upload.addEventListener('reload', this._get);
        this.control.addEventListener('click', this._toggleGrid);
        setInterval(this._get, 20000)
    }

    _get (event) {
        return fetch('/files')
            .then(res => res.json())
            .then(files => {
                if (!this._files) {
                    this._renderFiles(files);
                }

                if (this._files.length != files.length) {
                    this._renderFiles(files);
                }
            })
    }

    _renderFiles (files) {
        this._files = files;

        this.element.innerHTML = '';

        for (let file of files) {
            let el = document.createElement('div')

            let name;

            if (file.fileName.length > 13) {
                name = `${file.fileName.substring(0, 10)}...`;
            } else {
                name = file.fileName;
            }

            el.classList.add('file')
            el.innerHTML = `
                <img src="/static/img/file.svg" alt="File" width="20px">
                <span>${name}</span>
            `

            el.addEventListener('click', this._onClick);

            this.element.appendChild(el);
        }
    }

    _onClick (event) {
        document.querySelectorAll('.selected').forEach(selected => {
            selected.classList.remove('selected');
        })
        event.target.classList.add('selected');
    }

    _toggleGrid () {
        this.control.classList.add('spin')
        this.control.addEventListener('animationend', () => this.control.classList.remove('spin'));
        if (this.element.classList.contains('columns')) {
            this.element.classList.add('grid');
            this.element.classList.remove('columns');
        } else {
            this.element.classList.add('columns');
            this.element.classList.remove('grid');
        }
    }
}