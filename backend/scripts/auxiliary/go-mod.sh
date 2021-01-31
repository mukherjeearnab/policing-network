cd ../../chaincode

CC_DIR=$PWD

CC_NAMES="chargesheet_cc citizenprofile_cc evidence_cc fir_cc investigation_cc judgement_cc"

for CC in $CC_NAMES; do
    echo "Installing Go dependencies in "$CC
    cd $CC
    go mod vendor
    cd ..
done
echo "Installing Go dependencies complete!"
