.SILENT:
run:
	${info RUNNING ${DAY}: ${QUESTION} -}
	${info }
	go build -o solver ./src/${DAY}/${QUESTION}
	./solver ./src/${DAY}/input.txt
	rm solver
