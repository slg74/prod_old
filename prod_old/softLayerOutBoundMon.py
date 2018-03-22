#!/usr/bin/python

import csv
import operator
import os
import smtplib
import sys
import SoftLayer

from email.MIMEMultipart import MIMEMultipart
from email.MIMEText import MIMEText
from email.MIMEBase import MIMEBase
from email import encoders

total_out = []

def getPublicBandwidth(THRESHHOLD):

    client = SoftLayer.Client()
    theMask = "mask[outboundPublicBandwidthUsage]"
    result = client['SoftLayer_Account'].getHardware()

    f = open('public_outbound.csv','w')
    for server in result:
        serverInfo = client['SoftLayer_Hardware_Server'].getObject(id=server['id'],mask=theMask)
        pubout = float(serverInfo.get('outboundPublicBandwidthUsage',0.0))
        name = serverInfo['fullyQualifiedDomainName']

        if pubout > THRESHHOLD:
            total_out.append(pubout)
            print(name + "," + str(pubout))
            s = name + "," + str(pubout) + "\n"
            f.write(s)

    f.close()


def getCost():

    price_per_gb_overage = 0.09
    s = "total: " + str(sum(total_out)) + "GB ...  price: $" + str(round(sum(total_out)*price_per_gb_overage,2))
    print s
    return s


def sortCsv():

    data = csv.reader(open('public_outbound.csv'), delimiter=',')
    sortedlist = sorted(data, key=lambda x: float(x[1]), reverse=True)

    with open('public_outbound_sorted.csv','wb') as f:
        fileWriter = csv.writer(f, delimiter=',')
        for row in sortedlist:
            fileWriter.writerow(row)


def getSMTPPassword():

    credentials = {}
    with open('.noreplypw','r') as f:
        for line in f:
            user, pw = line.strip().split(':')

    return pw


def emailSortedCsv(RECIPIENTS):

    fromaddr = 'noreply@r1soft.com'
    toaddr = RECIPIENTS

    msg = MIMEMultipart()

    msg['From'] = fromaddr
    msg['To'] = toaddr
    msg['Subject'] = 'SoftLayer - Daily Production Outbound Public Interface Bandwidth Usage Report'

    body = 'Current outbound public interface bandwidth usage report for prod env (see attached).'

    msg.attach(MIMEText(body, 'plain'))

    filename = 'public_outbound_sorted.csv'

    # add cost to csv file
    cost = getCost()

    # append to csv
    with open(filename, 'ab') as f:
        f.write(cost)

    attachment = open(filename, 'rb')

    part = MIMEBase('application', 'octet-stream')
    part.set_payload((attachment).read())
    encoders.encode_base64(part)
    part.add_header('Content-Disposition', 'attachment; filename= %s' % filename)
    msg.attach(part)

    pw = getSMTPPassword()

    server = smtplib.SMTP('smtp.office365.com', 587)
    server.starttls()
    server.login(fromaddr, pw) 
    text = msg.as_string()
    server.sendmail(fromaddr, toaddr, text)
    server.quit()
    
    f.close()
    attachment.close()


def main():

    #THRESHHOLD = 500.0
    THRESHHOLD = 0.0
    #RECIPIENTS = 'scott.gillespie@r1soft.com,alex.vongluck@r1soft.com,stan.love@r1soft.com,tariq.siddiqui@r1soft.com'
    #RECIPIENTS = 'scott.gillespie@r1soft.com,tim.parker@r1soft.com'

    getPublicBandwidth(THRESHHOLD)
    sortCsv()
    #emailSortedCsv(RECIPIENTS) 
    

if __name__ == "__main__":

    try:
        main()
    except KeyboardInterrupt:
        sys.exit()


