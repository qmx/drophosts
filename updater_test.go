package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateHostsWithEntries(t *testing.T) {
	original := `
#
# /etc/hosts: static lookup table for host names
#

#<ip-address>   <hostname.domain.org>   <hostname>
127.0.0.1       localhost.localdomain   localhost
::1             localhost.localdomain   localhost
10.144.95.3     somerandomhost

# End of file

## drophosts ##
1.2.3.4 hostA.kubelocal
2.3.4.5 hostB.kubelocal
## drophosts ##`
	newContent := `## drophosts ##
3.4.5.6 hostA.kubelocal
4.5.6.7 hostB.kubelocal
## drophosts ##`
	expectedOutput := `
#
# /etc/hosts: static lookup table for host names
#

#<ip-address>   <hostname.domain.org>   <hostname>
127.0.0.1       localhost.localdomain   localhost
::1             localhost.localdomain   localhost
10.144.95.3     somerandomhost

# End of file

## drophosts ##
3.4.5.6 hostA.kubelocal
4.5.6.7 hostB.kubelocal
## drophosts ##`
	assert.Equal(t, expectedOutput, UpdateHosts(original, newContent))
}

func TestUpdateHostsWithNoEntries(t *testing.T) {
	original := `
#
# /etc/hosts: static lookup table for host names
#

#<ip-address>   <hostname.domain.org>   <hostname>
127.0.0.1       localhost.localdomain   localhost
::1             localhost.localdomain   localhost
10.144.95.3     somerandomhost

# End of file`
	newContent := `
## drophosts ##
1.2.3.4 hostA.kubelocal
2.3.4.5 hostB.kubelocal
## drophosts ##`

	expectedOutput := `
#
# /etc/hosts: static lookup table for host names
#

#<ip-address>   <hostname.domain.org>   <hostname>
127.0.0.1       localhost.localdomain   localhost
::1             localhost.localdomain   localhost
10.144.95.3     somerandomhost

# End of file

## drophosts ##
1.2.3.4 hostA.kubelocal
2.3.4.5 hostB.kubelocal
## drophosts ##`
	assert.Equal(t, expectedOutput, UpdateHosts(original, newContent))
}
