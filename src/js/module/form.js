'use strict';

document.addEventListener('DOMContentLoaded', () => {
  const inputs = document.getElementsByTagName('input');
  const form = document.forms.namedItem('article-form');
  const saveBtn = document.querySelector('.article-form__save');
  const cancelBtn = document.querySelector('.article-form__cancel');
  const previewOpenBtn = document.querySelector('.article-form__open-preview');
  const previewCloseBtn = document.querySelector('.article-form__close-preview');
  const articleFormBody = document.querySelector('.article-form__body');
  const articleFormPreview = document.querySelector('.article-form__preview');
  const articleFormBodyTextArea = document.querySelector('.article-form__input--body');
  const articleFormPreviewTextArea = document.querySelector('.article-form__preview-body-contents');
  const errors = document.querySelector('.article-form__errors');
  const errorTmpl = document.querySelector('.article-form__error-tmpl').firstElementChild;
  const csrfToken = document.getElementsByName('csrf')[0].content;

  const mode = { method: '', url: '' };
  if (window.location.pathname.endsWith('new')) {
    mode.method = 'POST';
    mode.url = '/articles';
  } else if (window.location.pathname.endsWith('edit')) {
    mode.method = 'PATCH';
    mode.url = `/articles/${window.location.pathname.split('/')[2]}`;
}
  const { method, url } = mode;
  for (let elm of inputs) {
    elm.addEventListener('keydown', event => {
      if (event.keyCode && event.keyCode === 13) {
        event.preventDefault();
        return false;
      }
    });
  }

  previewOpenBtn.addEventListener('click', event => {
    articleFormPreviewTextArea.innerHTML = md.render(articleFormBodyTextArea.value);

    articleFormBody.style.display = 'none';

    articleFormPreview.style.display = 'grid';
  });

  previewCloseBtn.addEventListener('click', event => {
    articleFormBody.style.display = 'grid';

    articleFormPreview.style.display = 'none';
  });

  cancelBtn.addEventListener('click', event => {
    event.preventDefault();

    window.location.href = url;
  });

  saveBtn.addEventListener('click', event => {
    event.preventDefault();

    errors.innerHTML = null;

    const fd = new FormData(form);

    let status;

    fetch(`/api${url}`, {
      method: method,
      headers: { 'X-CSRF-Token': csrfToken },
      body: fd
    })
      .then(res => {
        status = res.status;
        return res.json();
      })
      .then(body => {
        console.log(JSON.stringify(body));

        if (status === 200) {
          window.location.href = url;
        }

        if (body.ValidationErrors) {
          showErrors(body.ValidationErrors);
        }
      })
      .catch(err => console.error(err));
  });
  const showErrors = messages => {
    if (Array.isArray(messages) && messages.length != 0) {
      const fragment = document.createDocumentFragment();

      messages.forEach(message => {
        const frag = document.createDocumentFragment();

        frag.appendChild(errorTmpl.cloneNode(true));

        frag.querySelector('.article-form__error').innerHTML = message;

        fragment.appendChild(frag);
      });

      errors.appendChild(fragment);
    }
  };
});
