pkgname=cutie
pkgver=0.1.0
pkgrel=1
pkgdesc="A cute CLI project manager to organize your projects :33"
arch=('any')
url="https://github.com/goferwplynie/cutie"
license=('MIT')
depends=('go')
source=("git+https://github.com/goferwplynie/cutie.git")
sha256sums=('SKIP')

build() {
    cd "$srcdir/cutie"
    go build -o cutie
}

package() {
    install -Dm755 "$srcdir/cutie/cutie" "$pkgdir/usr/bin/cutie"
}

