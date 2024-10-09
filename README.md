# Resizem


**Resizem (combine "resize" with "them")** is an app designed for **bulk image process.**  It is particularly useful for users who need to resize, convert, and manage large numbers of image files at once. It supports a variety of formats like JPG, JPEG, PNG, GIF, BMP, TIFF and WebP. It also allows you to set custom dimensions while ensuring the quality of the image remains intact (with resampling filters).

![Screenshot](screenshot.png)

<p align="center">
<a target="_blank" href="https://github.com/barats/resizem/stargazers"><img src="https://img.shields.io/github/stars/barats/resizem"/></a> 
&nbsp;&nbsp;&nbsp;
<a target="_blank" href="https://gitee.com/barat/resizem/stargazers"><img src="https://gitee.com/barat/resizem/badge/star.svg?theme=dark"/></a>
</p>

## Features

1. Bulk image reszie and convert
1. Simple & easy UI with support for custom settings
1. Rich format support: JPG,JPEG,PNG,GIF,BMP,TIF,TIFF and WebP
1. Multiple resamping filters: NearestNeighbor, Box, Linear, Hermite, MitchellNetravali, CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine  

## Technical details

**Resizem** uses **Golang** to do the "resize and convert" work and **Svelte** + **Flowbite Svelte** + **TailwindCSS** for frontend work.  
 
**This is glued together as a single binary with native rendering by the fantastic [Wails](https://wails.io) framework.**

## License 

```
Copyright (c) [2024] [Barat Semet]
[Resizem] is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
```

## Give Thanks To 

1. [Wails.io](https://wails.io) 
1. [Flowbite Svelte](https://flowbite-svelte.com)
1. [TailwindCSS](https://tailwindcss.com)
1. [disintegration/imaging](https://github.com/disintegration/imaging)

## Collaboration

If you find this app useful, or are interested in open source collaboration, you can try:

1. Actively provide feedback on problems encountered while you're using this app
1. Read source code and submit appropriate and necessary PRs
1. Translate this app into other languages(put translated language files in `frontend/src/lib/i18n/locales/` and draft new PR)


## Release

1. Github [https://github.com/barats/resizem](https://github.com/barats/resizem)
1. Gitee [https://gitee.com/barat/resizem](https://gitee.com/barat/resizem)

## Github Stargazers
[![Stargazers over time](https://starchart.cc/barats/resizem.svg?variant=adaptive)](https://starchart.cc/barats/resizem)