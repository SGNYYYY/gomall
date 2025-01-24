.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/SGNYYYY/gomall/app/frontend