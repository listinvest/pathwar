# this file is generated by 'make generate-register'
pathwar admin challenge-add --slug training-sqli --name 'Training: Sqli' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/training-sqli --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge training-sqli --body 'Learn to execute custom code.' --passphrases 2 --category web --redump-policy '[{"strategy":"every","delay":"1d"}]' --tags tutorial,sql,injection --validation-reward 5 --compose-bundle ./web/training-sqli//pathwar-compose.yml
pathwar admin season-challenge-add --flavor training-sqli --season global
pathwar admin challenge-add --slug training-brute --name 'Training: Brute' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/training-brute --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge training-brute --body 'Learn to brute force.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags tutorial,bruteforce --validation-reward 5 --compose-bundle ./web/training-brute//pathwar-compose.yml
pathwar admin season-challenge-add --flavor training-brute --season global
pathwar admin challenge-add --slug captcha-mario --name Mario --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/captcha-mario --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge captcha-mario --body 'Bypass this Captcha protection.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags bruteforce --validation-reward 15 --purchase-price 5 --compose-bundle ./web/captcha-mario//pathwar-compose.yml
pathwar admin season-challenge-add --flavor captcha-mario --season global
pathwar admin challenge-add --slug helloworld --name 'Hello World!' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/helloworld --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge helloworld --body 'Introduction challenge.' --passphrases 1 --category web --redump-policy '[{"strategy":"every","delay":"1d"}]' --tags intro --validation-reward 1 --compose-bundle ./web/helloworld//pathwar-compose.yml
pathwar admin season-challenge-add --flavor helloworld --season global
pathwar admin challenge-add --slug training-include --name 'Training: Include' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/training-include --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge training-include --body 'Learn to execute custom code.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags tutorial,exec --validation-reward 5 --compose-bundle ./web/training-include//pathwar-compose.yml
pathwar admin season-challenge-add --flavor training-include --season global
pathwar admin challenge-add --slug training-http --name 'Training: HTTP' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/training-http --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge training-http --body 'Learn to play with HTTP.' --passphrases 5 --category web --redump-policy '[{"strategy":"every","delay":"1d"}]' --tags tutorial,http --validation-reward 5 --compose-bundle ./web/training-http//pathwar-compose.yml
pathwar admin season-challenge-add --flavor training-http --season global
pathwar admin challenge-add --slug upload-hi --name 'Upload Hi!' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/upload-hi --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge upload-hi --body 'Execute your code.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags exec --validation-reward 15 --purchase-price 5 --compose-bundle ./web/upload-hi//pathwar-compose.yml
pathwar admin season-challenge-add --flavor upload-hi --season global
pathwar admin challenge-add --slug imageboard --name 'Image Board' --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/imageboard --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge imageboard --body 'The passphrase is commented somewhere in the source.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags injection --validation-reward 15 --purchase-price 5 --compose-bundle ./web/imageboard//pathwar-compose.yml
pathwar admin season-challenge-add --flavor imageboard --season global
pathwar admin challenge-add --slug captcha-luigi --name Luigi --homepage https://github.com/pathwar/pathwar/tree/master/challenges/web/captcha-luigi --locale en --author 'Pathwar Staff'
pathwar admin challenge-flavor-add --challenge captcha-luigi --body 'Bypass this Captcha protection.' --passphrases 1 --category web --redump-policy '[{"strategy":"on-validation"},{"strategy":"every","delay":"1d"}]' --tags bruteforce --validation-reward 15 --purchase-price 5 --compose-bundle ./web/captcha-luigi//pathwar-compose.yml
pathwar admin season-challenge-add --flavor captcha-luigi --season global
