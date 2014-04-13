PIFILE=/home/ted/pi/workspace/2014-01-07-wheezy-raspbian.img
PIMNT=/home/ted/pi/workspace/pimnt
LOOPDEV=`losetup -f`
VNW=/home/ted/src/vnw
PIMNT=/

pimount()
{
losetup $LOOPDEV $PIFILE
partx -a $LOOPDEV
mount ${LOOPDEV}p2 $PIMNT
}

piumount()
{
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
# pimount
#Stuff to build and deploy master program.
cp ~/dev/libnfc-1.7.1.tar.bz2 $PIMNT/home/pi/libnfc.tar.bz2
cd $PIMNT/home/pi/
tar -xf libnfc.tar.bz2
# Copy over main function
mkdir -p $PIMNT/home/pi/src/
install -D -o ted $VNW $PIMNT/home/pi/src/

# Copy important scripts
install $VNW/scripts/firstrun $PIMNT/etc/init.d/
ln -s /etc/init.d/firstrun $PIMNT/etc/rc2.d/S90firstrun
install -d $PIMNT/service/
install -o ted $VNW/scripts/run $PIMNT/service/main/run
#piumount
# This is close, but not necessarially right. cigar; hrmm.
# truncate -s 4008706048
}

vnw
