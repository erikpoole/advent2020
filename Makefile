.SILENT:
run:
	${info RUNNING ${DAY}: ${QUESTION} -}
	${info }
	go build -o solver ./${DAY}/${QUESTION}
	./solver ${DAY}/input.txt
	rm solver
