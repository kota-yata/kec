require "formula"

class Kec < Formula
  desc "Image format extension converter. Only for jpg/png/gif image. Poor"
  homepage "https://github.com/kota-yata/KEC"
  url "https://github.com/kota-yata/KEC/archive/0.0.1.tar.gz"
  sha256 "a6dbf920948a1da8a6a65aff35ea306b31a069b1134e68a401e411f99983f7be"
  license ""

  # depends_on "cmake" => :build

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    system "./configure", "--disable-debug",
                          "--disable-dependency-tracking",
                          "--disable-silent-rules",
                          "--prefix=#{prefix}"
    # system "cmake", ".", *std_cmake_args
  end

  # test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test KEC`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    #system "false"
  # end
end
