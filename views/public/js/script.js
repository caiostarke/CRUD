var cardBack = document.querySelector('.card__back')
var cardFront = document.querySelector('.card__front')

isFront = true;
if (isFront) {
    cardBack.style.display = 'none';
}

function flip() {
    if (isFront) {
        cardFront.style.display = 'none';
        cardBack.style.display = 'flex';
        isFront = false;
    } else {
        cardBack.style.display = 'none';
        cardFront.style.display = 'flex';
        isFront = true;
    }
}