update:
	chmod +x hz_gen.sh \
	&& ./hz_gen.sh
run:
	chmod +x build.sh \
	&& ./build.sh \
	&& ./output/bootstrap.sh