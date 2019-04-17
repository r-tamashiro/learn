# coding: UTF-8
import urllib2
from bs4 import BeautifulSoup

# アクセスするURL
url = "https://www.google.co.jp/"
url = "https://ja.wikipedia.org/wiki/Linux"

html = urllib2.urlopen(url)

soup = BeautifulSoup(html, "html.parser")

title_tag = soup.title

title = title_tag.string

print title_tag

print title
