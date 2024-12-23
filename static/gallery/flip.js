$(document).on('click', '.gallery-component > div', function(e) {
    if (!$(e.target).closest('button, input, select, a, pre').length) {
        $(this).toggleClass('flipped');
    }
});