package main

import (
    "os"
    "io/ioutil"
    "image"
    "image/png"
    "math"
    "bytes"
    "fmt"
    "encoding/hex"
)


func main() {
    // Check command line arguments
    if len(os.Args) != 3 {
        fmt.Println("Usage: js2png <js_file> <png_file>")
        os.Exit(1)
    }

    var src_filename = os.Args[1]
    var dst_filename = os.Args[2]

    // Read source file
    src, err := ioutil.ReadFile(src_filename)
    if err != nil { panic(err) }

    // Create data buffer
    var buffer = new(bytes.Buffer)

    // Write payload size
    var payload_sz = len(src)

    buffer.WriteByte(0)
    buffer.WriteByte(uint8(payload_sz >> 56 & 0xff))
    buffer.WriteByte(uint8(payload_sz >> 48 & 0xff))
    buffer.WriteByte(255)
    buffer.WriteByte(uint8(payload_sz >> 40 & 0xff))
    buffer.WriteByte(uint8(payload_sz >> 32 & 0xff))
    buffer.WriteByte(uint8(payload_sz >> 24 & 0xff))
    buffer.WriteByte(255)
    buffer.WriteByte(uint8(payload_sz >> 16 & 0xff))
    buffer.WriteByte(uint8(payload_sz >> 8 & 0xff))
    buffer.WriteByte(uint8(payload_sz & 0xff))
    buffer.WriteByte(255)

    // Write payload into buffer as RGBA pixels
    var data_sz = int(math.Ceil(float64(payload_sz) / 3))

    for i := 0; i < data_sz; i++ {
        p := i * 3

        _, err := buffer.Write(src[p:p + 3])
        if err!= nil { panic(err) }

        buffer.WriteByte(255)
    }

    // Create bitmap to fit the payload
    var bitmap_sz = int(buffer.Len() / 4)
    w := int(math.Ceil(math.Sqrt(float64(bitmap_sz))))
    h := int(math.Ceil(float64(bitmap_sz) / float64(w)))

    rect := image.Rect(0, 0, w, h)
    img := image.NewNRGBA(rect)
    img.Pix = make([]uint8, w * h * 4)

    // Copy payload into bitmap
    copy(img.Pix, buffer.Bytes())

    // Log
    fmt.Printf("Source:\n%s\n", hex.Dump(src))
    fmt.Printf("Buffer:\n%s\n", hex.Dump(buffer.Bytes()))
    fmt.Printf("Image:\n%s\n", hex.Dump(img.Pix))

    // Write bitmap to PNG file
    dst, err := os.Create(dst_filename)
    if err != nil { panic(err) }
    defer dst.Close()

    png.Encode(dst, img)
}
