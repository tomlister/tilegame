for f in $(find ./ -name '*.wav'); do 
ffmpeg -i $f -acodec pcm_s16le -ar 44100 ${f%.wav}z.wav; 
rm $f;
mv ${f%.wav}z.wav $f;
done
