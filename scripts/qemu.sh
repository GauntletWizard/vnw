#!/bin/bash
PIFILE=/home/ted/pi/2014-01-07-wheezy-raspbian.img
PIMNT=/home/ted/pi/pimnt
LOOPDEV=`losetup -f`
VNW=/home/ted/code/vnw
GO=/home/ted/dev/go
WIRELESS=/home/ted/dev/rpi-wireless
LIBNFC=/home/ted/dev/libnfc-1.7.1.tar.bz2
WPA=$PIMNT/etc/wpa_supplicant/wpa_supplicant.conf

if [ "$EUID" -ne 0 ];
then
	echo "Must be run as root"
	exit 1
fi


pimount()
{
losetup $LOOPDEV $PIFILE
partx -a $LOOPDEV
mount ${LOOPDEV}p2 $PIMNT
}

piumount()
{
cd /
sync
sleep 1
umount $PIMNT
losetup -d `losetup -a |grep $PIFILE |cut -d: -f1`
}

# Follows instructions from http://xecdesign.com/qemu-emulating-raspberry-pi-the-easy-way/
mkemu ()
{
pimount
echo "#/usr/lib/arm-linux-gnueabihf/libcofi_rpi.so" > $PIMNT/etc/ld.so.preload
echo 'KERNEL=="sda", SYMLINK+="mmcblk0"' > $PIMNT/etc/udev/rules.d/90-qemu.rules
echo 'KERNEL=="sda?", SYMLINK+="mmcblk0p%n"' >> $PIMNT/etc/udev/rules.d/90-qemu.rules
echo 'KERNEL=="sda2", SYMLINK+="root"' >> $PIMNT/etc/udev/rules.d/90-qemu.rules
piumount
}

mkreal ()
{
pimount
echo "/usr/lib/arm-linux-gnueabihf/libcofi_rpi.so" > $PIMNT/etc/ld.so.preload
rm $PIMNT/etc/udev/rules.d/90-qemu.rules
piumount
}

vnw ()
{
pimount
#Stuff to build and deploy master program.
cp -r $GO $PIMNT/home/pi/
cp $LIBNFC $PIMNT/home/pi/libnfc.tar.bz2
cd $PIMNT/home/pi/
tar -xf libnfc.tar.bz2

# Drivers
KERNEL=3.10.25+
install -p -m 644 $WIRELESS/8188eu.ko $PIMNT/lib/modules/$KERNEL/kernel/drivers/net/wireless
install -D -p -m 644 $WIRELESS/rtl8188eufw.bin $PIMNT/lib/firmware/rtlwifi/rtl8188eufw.bin

# Copy over main function
install -o ted -d $PIMNT/home/pi/src/
cp -r $VNW $PIMNT/home/pi/src/
chown -R ted:ted $PIMNT/home/pi/src/

# Copy important scripts
install $VNW/scripts/firstrun $PIMNT/etc/rc.local
install -d $PIMNT/etc/service
install -o ted -d $PIMNT/etc/service/main
install -o ted -D $VNW/scripts/run $PIMNT/etc/service/main/run
rm $PIMNT/etc/init.d/mathkernel

# Set up secrets
install -o ted $VNW/secrets/secretfile $PIMNT/home/pi/
install -d -m 600 $VNW/secrets/wpa $PIMNT/etc/wicd/wireless-settings.conf

# Make things run.
install -o ted $VNW/scripts/run $PIMNT/service/main/run
install $VNW/scripts/logrotate $PIMNT/etc/logrotate.d/main

# Make me hate things less.
echo "XKBLAYOUT=\"us\"" >> $PIMNT/etc/default/keyboard
patch $PIMNT/etc/inittab $VNW/scripts/inittab.patch

# Clear state.
rm $VNW/var/opt/vnw-run
piumount
# This is close, but not necessarially right. cigar; hrmm.
# truncate -s 4008706048
}

case $1 in
	pimnt|pimount) pimount ;;
	piumnt|piumount) piumount ;;
	mkemu) mkemu ;;
	mkdreal) mkreal ;;
	vnw) vnw ;;
	*) vnw ; mkemu ;;
esac
