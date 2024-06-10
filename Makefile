update:
	chmod +x hz_gen.sh \
	&& ./hz_gen.sh
run:
	export MYSQL_HOST=localhost \
	MINIO_HOST=localhost \
	&& chmod +x build.sh \
	&& ./build.sh \
	&& ./output/bootstrap.sh