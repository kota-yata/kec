class Kec < Formula
  desc "Image format extension converter. Only for jpg/png/gif image. Poor"
  homepage "https://github.com/kota-yata/KEC"
  url "https://github.com/kota-yata/KEC/releases/download/0.0.1/kec"
  sha256 "6a7d482dc4341ff8c7b2f8cabc6ffa73937c045c6f7368538624087c5947a1ba"
  license ""

  def install
    bin.install "kec"
  end

  test do
    system "true"
  end
end
