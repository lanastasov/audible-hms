audible library list
  B00AUVOMFS: Napoleon Hill: The Master Key to Riches
audible download --asin B00AUVOMFS --aaxc --cover --cover-size 1215 --chapter
audible download --asin B00AUVOMFS --aaxc --cover --cover-size 1215 
./AAXtoMP3/AAXtoMP3 ./The_Master_Key_to_Riches.aaxc

create empty json file if line 4 is used

-------------------------------------------------------------------------------

convert mp4 to mp3
ffmpeg -i video.mp4 -b:a 192K -vn music.mp3

-------------------------------------------------------------------------------

This will concatenate two mp3 files, and the resulting metadata will be that of the first file:

ffmpeg -i "concat:file1.mp3|file2.mp3" -acodec copy output.mp3

-------------------------------------------------------------------------------

#syncrhonize folders locally
rsync -cavu /home/lanastasov/Downloads/audio-books/ /run/media/lanastasov/New Volume/audio-books/

-------------------------------------------------------------------------------

#replace : as - in all filenames in current directory
for file in ./*; do
	mv "$file" "${file//:/-}"
done

-------------------------------------------------------------------------------

