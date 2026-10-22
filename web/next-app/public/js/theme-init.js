(function () {
  try {
    var t = localStorage.getItem('orchid-theme');
    var d;
    if (t === 'dark' || t === 'light') {
      d = t === 'dark';
    } else {
      d = window.matchMedia('(prefers-color-scheme:dark)').matches;
    }
    if (d) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  } catch (e) {}
})();
