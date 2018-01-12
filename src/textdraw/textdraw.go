/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

 package textdraw

 import(
     "github.com/veandco/go-sdl2/sdl"
 )

 type Window struct {
     renderer *sdl.Renderer
     width int32
     height int32
 }