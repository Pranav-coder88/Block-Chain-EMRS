// Footer React Component

import {ButtonGroup, Container, IconButton, Stack, Text} from '@chakra-ui/react';
import {FaGithub, FaLinkedin, FaTwitter} from 'react-icons/fa';
import {Logo} from './Logo';
import * as React from 'react';


export const Footer = () => (
    <Container as='footer' role='contentinfo' py={{base: '12', md: '5'}}>
        <Stack spacing={{base: '4', md: '5'}}>
            <Stack justify='space-between' direction='row' align='center'>
                <Logo/>
                <ButtonGroup variant='ghost'>
                    <IconButton
                        as='a'
                        href='#'
                        aria-label='LinkedIn'
                        icon={<FaLinkedin fontSize='1.25rem'/>}
                    />
                    <IconButton as='a' href='#' aria-label='GitHub' icon={<FaGithub fontSize='1.25rem'/>}/>
                    <IconButton
                        as='a'
                        href='#'
                        aria-label='Twitter'
                        icon={<FaTwitter fontSize='1.25rem'/>}
                    />
                </ButtonGroup>
            </Stack>
            <Text fontSize='sm' color='subtle'>
                &copy; {new Date().getFullYear()} Virtualica
            </Text>
        </Stack>
    </Container>
);


export default Footer;
