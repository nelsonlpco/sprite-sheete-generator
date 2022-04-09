# Sprite Sheete generator v1.0

### A simple sprite sheet generator will create png while preserving the opacity layer.

####
inform the directory containing the images for creating the sprite sheet, it will be sorted in alphabetical order, ensure that this rule exists in the file name

#### options

| Parameter | type   | default | Action | required |
|-----------|--------|---------|--------|----------|
| -p        | string |         | sprites resource path | yes |
| -a        | number |  2      | Interpolation algorithm<br/>1 - Nearest-neighbor<br/>2 - Bilinear<br>3 - Bicubic<br/>4 - Mitchell-Netravali<br/>5 - Lanczos(2)<br/>6 - Lanczos(3) |
| -s        | number |  1024   | Sprite sheet size |
| -0        | string |  ./spriteSheet.png | Sprite sheet size |

#### input sample

![input](./images/inputsample.png)


#### output sample

![spritesheet](./images/spriteSheet.png)