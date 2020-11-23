require 'formula'

class Kec < Formula
  desc "Image format extension converter. Only for jpg/png/gif image. Poor"
  homepage "https://github.com/kota-yata/kec"
  url "https://github.com/kota-yata/kec/archive/1.0.0.tar.gz"
  sha256 "9a569bae71ff0ce6a9fdc6e7e68a878480ed023997c5bd8ab3e54b5a22c46022"
  license "MIT"

  def install
    bin.install "kec"
  end

  test do
    system "true"
  end
end
