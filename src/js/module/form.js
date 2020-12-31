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

  const mode = { method: '', url: '' };
  if (window.location.pathname.endsWith('new')) {
    mode.method = 'POST';
    mode.url = '/';
  } else if (window.location.pathname.endsWith('edit')) {
    mode.method = 'PATCH';
    mode.url = `/${window.location.pathname.split('/')[1]}`;
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
});
