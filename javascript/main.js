"use strict";

const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

ctx.fill('evenodd')

const background_image = new Image()
background_image.src = '/javascript/assets/environment/back.png'

const music = new Audio('/javascript/assets/sound/platformer_level03_loop.ogg')

document.addEventListener('DOMContentLoaded', () => {
    background_image.addEventListener('load', () => {
        ctx.drawImage(background_image, 0, 0, canvas.clientWidth, canvas.clientHeight, 0, 0, canvas.clientWidth, canvas.clientHeight)
    })

    music.play()
})