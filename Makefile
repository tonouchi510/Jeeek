APP_NAME := jeeekapi
REPO := github.com/tonouchi510/Jeeek


# goa周り
goagen:
	@goa gen ${REPO}/design

example:
	@goa example $(REPO)/design
	@script/fix_goa_example.sh

clean:
	@rm -rf cmd/
	@rm -rf gen/
	@rm *.go
